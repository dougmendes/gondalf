Feature: Provide all attributes about a ec2

    Scenario:List all ec2 from an aws account
        Given a region <instance_region>
        When the user request all ec2 from that region
        Then respond with a list of all ec2