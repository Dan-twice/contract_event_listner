// SPDX-License-Identifier: MIT

pragma solidity ^0.8.3;

contract Receiver {
    string public testReturn = "Works!";

    event TestEvent(string success);

    function TestEvents() public {
        emit TestEvent("Events listner works!");
    }

    // index keyword allow to search by; relevant only to logged event
    event ValueSender(
        address from,
        uint256 money
    );

    receive() external payable {
        emit ValueSender(msg.sender, msg.value);
    }
}
