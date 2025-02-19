import os
import requests
import asyncio
from dotenv import load_dotenv
from nlp import interpret_message
from aiogram import Bot, Dispatcher, types

load_dotenv()

TOKEN = os.getenv('TOKEN_TELEGRAM')
API_URL = "http://localhost:8080/transacoes"

dp = Dispatcher()

@dp.message()
async def process_message(message: types.Message):
    # await message.reply(f"Mensagem recebida")
    amount, category, type, date, payment, description = interpret_message(message.text)
    
    if description and category and amount and type and payment and date:
        data = {
            "description": description,
            "category": category,
            "amount": amount,
            "typ": type,
            "payment": payment,
            "obs": "data from bot telegram",
            "date": date,
        }
        print(data)

        response = requests.post(API_URL, json=data)

        if response.status_code == 201:
            await message.reply(f"Transação registrada: R${amount} em {description} ({category}) no dia {date}.")
        else:
            await message.reply("Erro ao registrar transação.")
    else:
        await message.reply("Não consegui entender. Tente começar pela data exemplo: 'Ontem gastei 150 no mercado'.")

async def main():
    bot = Bot(token=TOKEN)
    await dp.start_polling(bot, skip_updates=True)

if __name__ == "__main__":
    asyncio.run(main())