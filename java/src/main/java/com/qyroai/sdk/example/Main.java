package com.qyroai.sdk.example;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

import com.qyroai.sdk.QyroClient;
import com.qyroai.sdk.QyroServerClient;
import com.qyroai.sdk.auth.ClientTokenGenerator;
import com.qyroai.sdk.models.Message;
import com.qyroai.sdk.models.Session;

public class Main {

    private static final String BASE_URL = "https://qyroai.com";
    private static final String API_KEY_ID = "<>";
    private static final String API_KEY_SECRET = "<>";
    private static final String ASSISTANT_ID = "<>";

    public static void main(String[] args) {
        // ---------------------------
        // Using Server Client
        // ---------------------------
        QyroServerClient serverClient = new QyroServerClient(BASE_URL, API_KEY_ID, API_KEY_SECRET, 30);

        System.out.println("👉 Creating session via server API...");
        Map<String, Object> context = new HashMap<>();
        context.put("user_id", "user_abc");
        context.put("context", "first test run");

        Session session = serverClient.createSession(ASSISTANT_ID, context);
        System.out.println("✅ Session created: " + session.getId());

        System.out.println("👉 Sending message to assistant (server)...");
        List<Message> messages = serverClient.chat(ASSISTANT_ID, session.getId(), "Hello from Java Server Client!");
        for (Message m : messages) {
            System.out.println("[" + m.getRole() + "] " + m.getContent());
        }

        // ---------------------------
        // Using Client SDK (token-based)
        // ---------------------------
        ClientTokenGenerator tokenGen = new ClientTokenGenerator(API_KEY_ID, API_KEY_SECRET);
        String clientToken = tokenGen.generate(Map.of("user_id", "123"));

        QyroClient client = new QyroClient(BASE_URL, clientToken, 30);

        System.out.println("👉 Creating session via client API...");
        Session clientSession = client.createSession(ASSISTANT_ID, Map.of("context", "from client SDK"));
        System.out.println("✅ Client session created: " + clientSession.getId());

        System.out.println("👉 Sending message to assistant (client)...");
        List<Message> clientMessages = client.chat(ASSISTANT_ID, clientSession.getId(), "Hello from Java Client SDK!");
        for (Message m : clientMessages) {
            System.out.println("[" + m.getRole() + "] " + m.getContent());
        }
    }
}
