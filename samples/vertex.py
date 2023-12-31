from vertexai.preview.language_models import TextGenerationModel
from google.cloud import aiplatform
from google import cloud


def interview(temperature: float = .2):
    """Ideation example with a Large Language Model"""

    # TODO developer - override these parameters as needed:
    parameters = {
        "temperature": temperature,
        "max_output_tokens": 256,
        "top_p": .8,
        "top_k": 40,
    }

    model = TextGenerationModel.from_pretrained("text-bison@001")
    response = model.predict(
        'Give me ten interview questions for the role of program manager.',
        **parameters,
    )
    print(f"Response from Model: {response.text}")

interview()