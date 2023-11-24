import openai
client = openai.OpenAI(base_url="http://localhost:9999/v1")
# client = openai.OpenAI()

response = client.chat.completions.create(
    model="gpt-4-1106-preview",
    messages=[
        {
            "role": "user",
            "content": [
                {"type": "text", "text": "Tell me a joke about AI? Return in a JSON format with two keys 'joke' and 'category' "},
            ],
        }
    ],
    max_tokens=300,
    seed=1234,
    response_format={ "type": "json_object" }
)

print(response)