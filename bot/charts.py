import requests
import matplotlib.pyplot as plt
import pandas as pd

API_URL = "http://localhost:8080/transacoes"

# Get transactions
def get_transactions():
    response = requests.get(API_URL)
    return pd.DataFrame(response.json())

# Generate a pie chart
def generate_pie_chart(df):
    plt.figure(figsize=(6, 6))
    df.groupby("category")["amount"].sum().plot(kind="pie", autopct="%1.1f%%")
    plt.title("Gastos por Categoria")
    plt.ylabel("")
    plt.show()

# Generate a bar chart
def generate_bar_chart(df):
    df["date"] = pd.to_datetime(df["date"])
    df["month"] = df["date"].dt.strftime("%Y-%m")

    pivot_df = df.pivot_table(values="amount", index="month", columns="category", aggfunc="sum")

    plt.figure(figsize=(8, 5))
    pivot_df.plot(kind="bar", stacked=True, colormap="viridis", ax=plt.gca())
    plt.title("Comaparativo de Gastos por Categoria (Mensal)")
    plt.xlabel("Mês")
    plt.ylabel("Total Gasto")
    plt.xticks(rotation=45)
    plt.legend(title="Categoria", bbox_to_anchor=(1, 1))
    plt.grid(axis="y", linestyle="--", alpha=0.7)

    # plt.show()
    plt.savefig("bot/charts/bar_chart.png", bbox_inches="tight")
    plt.close()

# Test
if __name__ == "__main__":
    generate_bar_chart(get_transactions())