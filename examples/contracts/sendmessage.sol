pragma solidity ^0.4.0;

contract SendMessage {
    struct Message{ // Struct
        string text;
        string author;
    }

    Message[] public messages;

    function addMessage(string text, string author) {
        Message memory newMessage = Message(text, author);
        messages.push(newMessage);
    }

}
