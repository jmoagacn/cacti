/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

import { GluegunCommand } from 'gluegun'
import { fabricHelper, getKeyAndCertForRemoteRequestbyUserName } from '../../helpers/fabric-functions'
import logger from '../../helpers/logger'
import { Utils, ICryptoKey } from 'fabric-common'
import { commandHelp, getNetworkConfig, handlePromise } from '../../helpers/helpers'
import { EventsManager } from '@hyperledger-labs/weaver-fabric-interop-sdk'
import { EventSubscriptionState, EventType } from "@hyperledger-labs/weaver-protos-js/common/events_pb";
import * as fs from 'fs'
import * as path from 'path'

const command: GluegunCommand = {
  name: 'unsubscribe',
  alias: ['-sd'],
  description: 'Initiate event unsubscribe',
  run: async toolbox => {
    const {
      print,
      parameters: { options, array }
    } = toolbox
    if (options.help || options.h) {
      commandHelp(
        print,
        toolbox,
        `fabric-cli event unsubscribe --network=<network1|network2> --user=user1 --request-id="abc123" src/data/event_sub_sample.json`,
        'fabric-cli event unsubscribe --network=<network-name> --user=<user-id> --request-id=<request_id> <filename>',
        [
            {
                name: '--network',
                description:
                    'Local network for command. <network1|network2>'
            },
            {
                name: '--user',
                description:
                    'User for subscription. Default: user1'
            },
            {
                name: '--request-id',
                description:
                    'Request ID received during subscription.'
            },
            {
                name: '--debug',
                description:
                    'Shows debug logs when running. Disabled by default. To enable --debug=true'
            }
        ],
        command,
        ['event', 'unsubscribe']
      )
      return
    }
    console.log("Event Subscription")
    if (options.debug === 'true') {
        logger.level = 'debug'
        logger.debug('Debugging is enabled')
    }
    if (array.length != 1) {
        print.error('Not enough arguments supplied')
        return
    }
    if (!options['request-id'])
    {
      print.error('--request-id needs to be specified')
      return
    }
    if (!options['user']) {
        options['user'] = `user1`     //Default user
    }
    
    const networkName = options['network']
    const user = options['user']
    const requestId = options['request-id']
    
    const filepath = path.resolve(array[0])
    const data = JSON.parse(fs.readFileSync(filepath).toString())
    
    const netConfig = getNetworkConfig(networkName)
    if (!netConfig.connProfilePath || !netConfig.channelName || !netConfig.chaincode) {
        throw new Error(`No valid config entry found for ${networkName}`)
    }
    
    const { gateway, wallet, contract } = await fabricHelper({
        channel: netConfig.channelName,
        contractName: process.env.DEFAULT_CHAINCODE ? process.env.DEFAULT_CHAINCODE : 'interop',
        connProfilePath: netConfig.connProfilePath,
        networkName,
        mspId: netConfig.mspId,
        logger,
        discoveryEnabled: true,
        userString: user
    })
    
    const [keyCert, keyCertError] = await handlePromise(
      getKeyAndCertForRemoteRequestbyUserName(wallet, user)
    )
    if (keyCertError) {
      throw new Error(`Error getting key and cert ${keyCertError}`)
    }
    
    const eventMatcher = EventsManager.createEventMatcher(data.event_matcher)
    const eventPublicationSpec = EventsManager.createEventPublicationSpec(data.event_publication_spec)
    
    try {
        const response = await EventsManager.unsubscribeRemoteEvent(
            contract,
            eventMatcher,
            eventPublicationSpec,
            requestId,
            networkName,
            netConfig.mspId,
            netConfig.relayEndpoint,
            { address: data.view_address, Sign: true },
            keyCert
        )
        
        if (response.getStatus() == EventSubscriptionState.STATUS.UNSUBSCRIBED) {
            console.log("Event Unsubscription Success for requestId:", response.getRequestId())
        } else {
            console.log("Unknown error")
        }
    } catch(e) {
        console.log("Error: ", e.toString())
    }

    process.exit()
  }
}


module.exports = command
