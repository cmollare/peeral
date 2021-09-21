Feature: Core features

    Multiple scenarios to test basic capabilities of the server

    Background: First peer has been launched
      Given a peer is connected
      And has been for 60 seconds
      And I can get it's ID

    Scenario: Connect to a peer
      When my peer is started
      Then my peer can connect to it


