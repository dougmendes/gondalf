Feature: Provide all attributes about a instances

    Scenario:List all instances from an aws account
        Given a region <region>
        When the user request all instances from that region
        Then respond with a list of all instances