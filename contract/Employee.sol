//SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

contract Employee {
    struct User {
        bytes32 userEmail;
        uint256 userTime;
        uint256 index;
    }

    mapping(address => User) public listUsers;
    address[] public userIndex;

    event LogNewUser(
        address indexed userAddress,
        uint256 index,
        bytes32 userEmail,
        uint256 userTime
    );
    event LogUpdateUser(
        address indexed userAddress,
        uint256 index,
        bytes32 userEmail,
        uint256 userTime
    );
    event LogDeleteUser(address indexed userAddress, uint256 index);

    function isUser(address userAddress) public view returns (bool isIndeed) {
        require(userIndex.length != 0, "List user address is empty !");
        return (userIndex[listUsers[userAddress].index] == userAddress);
    }

    function insertUser(
        address userAddress,
        bytes32 userEmail,
        uint256 userTime
    ) public returns (uint256 index) {
        userIndex.push(userAddress);
        listUsers[userAddress].userEmail = userEmail;
        listUsers[userAddress].userTime = userTime;
        listUsers[userAddress].index = userIndex.length - 1;
        emit LogNewUser(
            userAddress,
            listUsers[userAddress].index,
            userEmail,
            userTime
        );
        return userIndex.length - 1;
    }

    function deleteUser(address userAddress) public returns (uint256 index) {
        require(isUser(userAddress), "User not found !");
        uint256 rowToDelete = listUsers[userAddress].index;
        address lastKey = userIndex[userIndex.length - 1];
        userIndex[rowToDelete] = lastKey;
        listUsers[userAddress] = listUsers[lastKey];
        delete listUsers[lastKey];
        userIndex.pop();
        emit LogDeleteUser(userAddress, rowToDelete);
        emit LogUpdateUser(
            lastKey,
            rowToDelete,
            listUsers[lastKey].userEmail,
            listUsers[lastKey].userTime
        );
        return rowToDelete;
    }

    function getUser(address userAddress)
    public
    view
    returns (
        bytes32 userEmail,
        uint256 userTime,
        uint256 index
    )
    {
        require(isUser(userAddress), "User not found !");
        return (
            listUsers[userAddress].userEmail,
            listUsers[userAddress].userTime,
            listUsers[userAddress].index
        );
    }

    function updateUserEmail(address userAddress, bytes32 userEmail)
    public
    returns (bool success)
    {
        require(isUser(userAddress), "User not found !");
        listUsers[userAddress].userEmail = userEmail;
        emit LogUpdateUser(
            userAddress,
            listUsers[userAddress].index,
            userEmail,
            listUsers[userAddress].userTime
        );
        return true;
    }

    function updateUserTime(address userAddress, uint256 userTime)
    public
    returns (bool success)
    {
        require(isUser(userAddress), "User not found !");
        listUsers[userAddress].userTime = userTime;
        emit LogUpdateUser(
            userAddress,
            listUsers[userAddress].index,
            listUsers[userAddress].userEmail,
            userTime
        );
        return true;
    }

    function getUserCount() public view returns (uint256 count) {
        return userIndex.length;
    }

    function getUserAtIndex(uint256 index)
    public
    view
    returns (address userAddress)
    {
        return userIndex[index];
    }

    function getAllUsers() public view returns (User[] memory)
    {
        require(userIndex.length > 0, "List of users is empty !!");
        User[] memory allUsers = new User[](userIndex.length);
        for (uint i = 0; i < userIndex.length; i++)
        {
            allUsers[i] = listUsers[userIndex[i]];
        }

        return allUsers;
    }
}
