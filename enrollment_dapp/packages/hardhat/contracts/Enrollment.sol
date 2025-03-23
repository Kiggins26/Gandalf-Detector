//SPDX-License-Identifier: MIT
pragma solidity >=0.8.0 <0.9.0;


import "hardhat/console.sol";

contract Enrollment{
    // State Variables
    address public immutable owner;
    mapping(address => string) public walletToDiscordNameMapping;
    uint256 public enrollmentPrice;

    // Constructor: Called once on contract deployment
    constructor(address _owner) {
        owner = _owner;
    }

    // Modifier: used to define a set of rules that must be met before or after a function is executed
    // Check the withdraw() function
    modifier isOwner() {
        // msg.sender: predefined variable that represents address of the account that called the current function
        require(msg.sender == owner, "Not the Owner");
        _;
    }

    function enrollAddressToGandalfTwoFA(string memory _discordName) public payable {
        console.log("Trying to set address '%s' for the discord Account %s",  msg.sender, _discordName);

        if (msg.value >= enrollmentPrice) {
            walletToDiscordNameMapping[msg.sender] = _discordName;
        } else {
            revert("payment did not meet requried price");
        }
    }

    function removeAddressFromGandalfTwoFA(address _addressToBeRemoved) public {
        require(msg.sender == _addressToBeRemoved, "Requested address can only be removed by the given address");
        delete walletToDiscordNameMapping[_addressToBeRemoved];
    }

    /**
     * Function that allows the owner to withdraw all the Ether in the contract
     * The function can only be called by the owner of the contract as defined by the isOwner modifier
     */
    function withdraw() public isOwner {
        (bool success, ) = owner.call{ value: address(this).balance }("");
        require(success, "Failed to send Ether");
    }


    function updateEnrollmentPrice(uint256 _newPrice) public isOwner {
        enrollmentPrice = _newPrice;
    }

    /**
     * Function that allows the contract to receive ETH
     **/
    receive() external payable {}
}
