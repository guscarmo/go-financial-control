import re
import dateparser

# List of possibles categories
CATEGORIES = ["mercado", "poupança", "cartão", "aluguel", "combustível", "lazer", "salário", "investimento", "saúde", "educação", "transporte", "vestuário", "outros", "assinatura", "alimentação"]

# List of possibles type "ganho"
TYPES_GANHO = ["ganho", "ganhei", "recebi"]

# List of possibles type "custo""
TYPES_CUSTO = ["gastei", "paguei", "comprei", "despesa", "custo", "gasto", "transferi"]

# List of possibles payment methods
PAYMENT = ["dinheiro", "crédito", "débito", "pix", "salário"]

def get_description(message, amount, categories, type, date, payments):
    """Extrai a descrição da mensagem."""
    words = message.split()

    for word in words:
        clean_word = word.lower().strip(",.!?")
        
        # Ignore float or int
        if re.match(r"^\d+([.,]\d+)?$", clean_word):
            continue

        # Ignore if date
        if dateparser.parse(clean_word, languages=["pt"]) is not None:
            continue

        # Ignore if word in exceptions
        if (
            clean_word in str(amount) or 
            clean_word in categories or 
            clean_word in type or 
            clean_word in date or 
            clean_word in payments or 
            clean_word in PAYMENT or 
            clean_word in TYPES_GANHO or 
            clean_word in TYPES_CUSTO or 
            clean_word in CATEGORIES
        ):
            continue
        
        # If word has more than 3 letters, return it
        if len(clean_word) > 3:
            return clean_word
    
    return None 

def interpret_message(message):
    """Extract fields from message."""
    
    message = message.lower()

    # Get amount from message
    amount = re.search(r"(?<![/\-])\b\d+(?:[.,]\d+)?\b(?![/\-])", message)
    amount = float(amount.group().replace(",", ".")) if amount else None

    # Indentify category from message
    category = next((cat for cat in CATEGORIES if cat in message.lower()), None)

    # Indentify type of transaction
    type = None
    if any(t in message for t in TYPES_GANHO):
        type = "ganho"
    elif any(t in message for t in TYPES_CUSTO):
        type = "custo"

    # Indentify payment method
    payment = next((p for p in PAYMENT if p in message.lower()), None)

    # Interpret relative date
    for w in message.split():
        date = dateparser.parse(w, languages=["pt"])
        if date:
            date = date.strftime("%Y-%m-%d")
            break
        else:
            date = None

    description = get_description(message, amount, category, type, date, payment)

    return amount, category, type, date, payment, description

# Test
if __name__ == "__main__":
    msg = "Ontem gastei 150 no crédito em combustível com o ford"
    print(interpret_message(msg))
