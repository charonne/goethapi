pragma solidity ^0.4.0;

contract SendMessage {

    address owner;

    function SendMessage() {
        owner = msg.sender;
    }

    struct Message{ // Struct
        string text;
        string author;
        string categories;
        string url;
        string hash;
    }

    Message[] public messages;

    function addMessage(string text, string author, string categories, string url, string hash) returns(uint length) {
        require(msg.sender == owner);
        Message memory newMessage = Message(text, author, categories, url, hash);
        return messages.push(newMessage);
    }

    function updateMessage(uint id, string text, string author, string categories, string url, string hash) {
        require(msg.sender == owner);
        Message memory newMessage = Message(text, author, categories, url, hash);
        messages[id] = newMessage;
    }

}
