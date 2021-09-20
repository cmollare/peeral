Feature: Discovery with IPFS DHT

    In order to join the network
    As a peer
    I should be able to find other peers on the network

    Scenario: Standard discovery
      Given a peer is connected
      And has been for 5 seconds
      When my peer is started
      Then it should start listening
      And it should find peers