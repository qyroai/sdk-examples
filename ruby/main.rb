require "qyro_sdk"

BASE_URL      = "https://qyroai.com"
API_KEY_ID    = "<>"
API_KEY_SECRET = "<>"
ASSISTANT_ID  = "<>"

if __FILE__ == $0
  # --- Server-side Example ---
  puts "=== Server SDK Example ==="
  server_client = QyroSDK::QyroServerClient.new(
    base_url: BASE_URL,
    api_key_id: API_KEY_ID,
    api_key_secret: API_KEY_SECRET
  )

  # create a new session
  session = server_client.create_session(ASSISTANT_ID, { user_id: "123" })
  puts "Created server session: #{session.id}"

  # send a chat message
  messages = server_client.chat(ASSISTANT_ID, session.id, "Hello from Ruby (server)")
  messages.each { |m| puts "[#{m.role}] #{m.content}" }

  # --- Client-side Example ---
  puts "\n=== Client SDK Example ==="
  token_generator = QyroSDK::ClientTokenGenerator.new(API_KEY_ID, API_KEY_SECRET)
  token = token_generator.generate({ client: "ruby-sdk-example" })

  client = QyroSDK::QyroClient.new(base_url: BASE_URL, token: token)
  client_session = client.create_session(ASSISTANT_ID, { locale: "en" })
  puts "Created client session: #{client_session.id}"

  client_messages = client.chat(ASSISTANT_ID, client_session.id, "Hello from Ruby (client)")
  client_messages.each { |m| puts "[#{m.role}] #{m.content}" }
end
