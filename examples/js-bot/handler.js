function handler(message) {
    console.log("test", JSON.stringify(message["Text"]));
    SendMessage(ctx, {
        ChatId: message["Chat"]["Id"],
        Text: "from js",
    });
}