import openai
client = openai.OpenAI(base_url="http://localhost:9999/v1")
# client = openai.OpenAI()

completion = client.chat.completions.create(
  model="gpt-4-vision-preview",
  messages=[
    {"role": "system", "content": "You are a helpful assistant."},
    {"role": "user", "content": "Tell me a joke about apples"}
  ]
)

print(completion.choices[0].message)