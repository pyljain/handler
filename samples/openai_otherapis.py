import openai
client = openai.OpenAI(base_url="http://localhost:9999/v1")

resp = client.images.generate(
  model="dall-e-3",
  prompt="A cute baby sea otter",
  n=1,
  size="1024x1024"
)

print(resp)