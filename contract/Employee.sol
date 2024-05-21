//SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

contract Employee {
    struct UserStruct {
        bytes32 userEmail;
        uint256 userTime;
        uint256 index;
    }

    mapping(address => UserStruct) private userStructs;
    address[] private userIndex;

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
        return (userIndex[userStructs[userAddress].index] == userAddress);
    }

    function insertUser(
        address userAddress,
        bytes32 userEmail,
        uint256 userTime
    ) public returns (uint256 index) {
        userIndex.push(userAddress);
        userStructs[userAddress].userEmail = userEmail;
        userStructs[userAddress].userTime = userTime;
        userStructs[userAddress].index = userIndex.length - 1;
        emit LogNewUser(
            userAddress,
            userStructs[userAddress].index,
            userEmail,
            userTime
        );
        return userIndex.length - 1;
    }

    function deleteUser(address userAddress) public returns (uint256 index) {
        require(isUser(userAddress), "User not found !");
        uint256 rowToDelete = userStructs[userAddress].index;
        address lastKey = userIndex[userIndex.length - 1];
        userIndex[rowToDelete] = lastKey;
        userStructs[userAddress] = userStructs[lastKey];
        delete userStructs[lastKey];
        userIndex.pop();
        emit LogDeleteUser(userAddress, rowToDelete);
        emit LogUpdateUser(
            lastKey,
            rowToDelete,
            userStructs[lastKey].userEmail,
            userStructs[lastKey].userTime
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
            userStructs[userAddress].userEmail,
            userStructs[userAddress].userTime,
            userStructs[userAddress].index
        );
    }

    function updateUserEmail(address userAddress, bytes32 userEmail)
        public
        returns (bool success)
    {
        require(isUser(userAddress), "User not found !");
        userStructs[userAddress].userEmail = userEmail;
        emit LogUpdateUser(
            userAddress,
            userStructs[userAddress].index,
            userEmail,
            userStructs[userAddress].userTime
        );
        return true;
    }

    function updateUserTime(address userAddress, uint256 userTime)
        public
        returns (bool success)
    {
        require(isUser(userAddress), "User not found !");
        userStructs[userAddress].userTime = userTime;
        emit LogUpdateUser(
            userAddress,
            userStructs[userAddress].index,
            userStructs[userAddress].userEmail,
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

    function getAllUsers() public view returns (UserStruct[] memory, address[] memory)
    {
      UserStruct[] memory allUsers = new UserStruct[](userIndex.length);
      for (uint i = 0; i < userIndex.length; i++) 
      {
        allUsers[i] = userStructs[userIndex[i]];
      }

      return (allUsers, userIndex);
    }
}
