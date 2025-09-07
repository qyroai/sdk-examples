from qyro_sdk.auth import ClientTokenGenerator
from qyro_sdk.client import QyroClient
from qyro_sdk.server import QyroServerClient

BASE_URL = "https://qyroai.com"
API_KEY_ID = "<>"
API_KEY_SECRET = "<>"
ASSISTANT_ID = "<>"

if __name__ == '__main__':
    server_client = QyroServerClient(
        base_url=BASE_URL,
        api_key_id=API_KEY_ID,
        api_key_secret=API_KEY_SECRET,
        timeout=120.0
    )

    session_id = server_client.create_session(ASSISTANT_ID, {"userId": "123"}).id

    output_messages = server_client.chat(assistant_id=ASSISTANT_ID, session_id=session_id, message="Hello, who are you?")
    print(output_messages)

    client_token_generator = ClientTokenGenerator(api_key_id=API_KEY_ID, api_key_secret=API_KEY_SECRET)
    client_token = client_token_generator.generate({
        "userId": "123"
    })

    client = QyroClient(base_url=BASE_URL, token=client_token)
    session_id = client.create_session(ASSISTANT_ID, {"userId": "123"}).id

    output_messages = client.chat(assistant_id=ASSISTANT_ID, session_id=session_id, message="Hello, who are you?")
    print(output_messages)

