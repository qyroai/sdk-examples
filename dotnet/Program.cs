using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using QyroSdk;

class Program
{
    static async Task Main()
    {
        const string BASE_URL = "https://qyroai.com";
        const string API_KEY_ID = "<>";
        const string API_KEY_SECRET = "<>";
        const string ASSISTANT_ID = "<>";

        // ---------------------------
        // Using Server Client
        // ---------------------------
        var serverClient = new QyroServerClient(BASE_URL, API_KEY_ID, API_KEY_SECRET);

        Console.WriteLine("👉 Creating session via server API...");
        var session = await serverClient.CreateSessionAsync(ASSISTANT_ID, new Dictionary<string, object>
        {
            { "user_id", "user_abc" },
            { "context", "first test run" }
        });
        Console.WriteLine($"✅ Session created: {session.Id}");

        Console.WriteLine("👉 Sending message to assistant (server)...");
        var messages = await serverClient.ChatAsync(ASSISTANT_ID, session.Id, "Hello from C# Server Client!");
        foreach (var m in messages)
        {
            Console.WriteLine($"[{m.Role}] {m.Content}");
        }

        // ---------------------------
        // Using Client SDK (token-based)
        // ---------------------------
        Console.WriteLine("\n👉 Generating client token...");
        var tokenGen = new ClientTokenGenerator(API_KEY_ID, API_KEY_SECRET);
        var clientToken = tokenGen.Generate(new Dictionary<string, object>
        {
            { "user_id", "123" }
        });
        Console.WriteLine("✅ Client token generated.");

        var client = new QyroClient(BASE_URL, clientToken);

        Console.WriteLine("👉 Creating session via client API...");
        var clientSession = await client.CreateSessionAsync(ASSISTANT_ID, new Dictionary<string, object>
        {
            { "context", "from client SDK" }
        });
        Console.WriteLine($"✅ Client session created: {clientSession.Id}");

        Console.WriteLine("👉 Sending message to assistant (client)...");
        var clientMessages = await client.ChatAsync(ASSISTANT_ID, clientSession.Id, "Hello from C# Client SDK!");
        foreach (var m in clientMessages)
        {
            Console.WriteLine($"[{m.Role}] {m.Content}");
        }
    }
}

