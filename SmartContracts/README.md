

There are two contracts used in MedRec:

- _Agent_ - stores information about an agent in MedRec, this could be a provider, patient, research group, family member of a patient, etc. If the Agent is a patient then the contract stores a list of their _Relationship_ contracts. Otherwise the Agent has a list of the different permissions it references. For example: in the case of a provider this is a list of permissions needed for various categories it provides, and in the case of a pharmacy this is a list of permissions it requests from patrons.
The _Agent_ contract also allows for a list of custodians to make operations on behalf of the agent. Each of the agent and all custodians can have their access rights selectively enabled or disabled. This allows parents to be custodians of minors and for married couples to turn on and off the ability for their spouse to make decisions on their behalf.

- _Agent Group_ - stores information about a group of agents in MedRec. Agent Groups cannot initiate relationships with other agents, they can only be the second party entered into one. An Agent Group allows an agent to give access to a set of other agents in one step, without explicitly giving permission to each group member. En example use case would allow users to give the members of an Emergency Room Agent Group access to their blood type.

- _Relationship_ - represents the relationship between two Agents, referred to as the patron and the provider. This contract manages the access permissions for data stored by the provider about the patron. The patron and provider are allowed to read all data, but other viewers need to be added.

- _AgentRegistry_ - stores information about signers. This contract provides an easy way to get the list of all signers and their names. It also retrieves the agent contract associated with a particular address. After a signer has been voted in via this contract they will then be added to the actual blockchain validator set. Current signers should be watching this contract for a new signer to be approved and then propose the new signer address via the clique API.


Make sure when testing to add at least 11 separate ethereum accounts to your blockchain provider. Each account will be assigned a role for the tests, pharmacy, family member, agent, etc.
