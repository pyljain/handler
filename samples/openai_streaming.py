import openai
client = openai.OpenAI(base_url="http://localhost:9999/v1")

completion = client.chat.completions.create(
  model="gpt-3.5-turbo",
  messages=[
    {"role": "system", "content": "You are a helpful assistant."},
    {"role": "user", "content": "Print numbers 1 to 10 seperated by new line!"}
  ],
  stream=True
)

for chunk in completion:
  print(chunk.choices[0].delta.content, end="")
