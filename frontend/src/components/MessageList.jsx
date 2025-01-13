import React from "react";
import Message from "./Message";

const MessageList = ({ message }) => {
  return (
    <div className="flex flex-col h-full space-y-4">
      {message.map((msg, index) => (
        <Message
          key={index}
          message={msg}
        />
      ))}
    </div>
  );
};

export default MessageList;
