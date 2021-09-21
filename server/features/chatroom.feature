Feature: Chat room feature

    In order to send messages
    As a peer
    I should be able to join a room and send messages

    Scenario: Join a room and send a message
      Given a peer is connected
      And has been for 5 seconds
      When my peer is started
      And is connected to first peer
      Then it should join room named customChatRoom and send a message hello world
      And other peer should receive a message hello world