// URL da API do backend
const API_URL = "http://localhost:8080/transacoes";

// Referências aos elementos HTML
const form = document.getElementById("transaction-form");
const transactionsTable = document.getElementById("transactions-table");

// Função para buscar e exibir transações
async function fetchTransactions() {
  try {
    const response = await fetch(API_URL);
    if (!response.ok) throw new Error("Erro ao buscar transações.");

    const transactions = await response.json();

    if (Array.isArray(transactions)) {
        transactions.forEach(transaction => {
            if (transaction.amount !== undefined) {
                console.log(`Descrição: ${transaction.description}`);
                console.log(`Valor: ${transaction.amount.toFixed(2)}`); // Aqui usamos toFixed corretamente
                displayTransactions(transactions);
            } else {
                console.warn("Transação sem valor:", transaction);
            }
        });
    } else {
        console.error("Formato inesperado de resposta da API:", transactions);
    }

    
  } catch (error) {
    console.error(error.message);
  }
}

// Função para exibir transações na tabela
function displayTransactions(transactions) {
  transactionsTable.innerHTML = ""; // Limpa a tabela

  transactions.forEach((transaction) => {
    const row = document.createElement("tr");

    row.innerHTML = `
      <td>${transaction.description}</td>
      <td>${transaction.amount.toFixed(2)}</td>
      <td>${new Date(transaction.data).toLocaleDateString()}</td>
    `;

    transactionsTable.appendChild(row);
  });
}

// Função para adicionar uma nova transação
async function addTransaction(transaction) {
  try {
    const response = await fetch(API_URL, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(transaction),
    });

    if (!response.ok) throw new Error("Erro ao adicionar transação.");
    fetchTransactions(); // Atualiza a lista após adicionar
  } catch (error) {
    console.error(error.message);
  }
}

// Lida com o envio do formulário
form.addEventListener("submit", (event) => {
  event.preventDefault();

  // Obtém os valores do formulário
  const description = document.getElementById("description").value;
  const amount = parseFloat(document.getElementById("amount").value);
  const date = document.getElementById("date").value;
  const category = document.getElementById("category").value;
  const type = document.getElementById("type").value;
  const payment = document.getElementById("payment").value;
  const obs = document.getElementById("obs").value;


  if (!description || !amount || !date || !category || !type || !payment || !obs ) {
    alert("Preencha todos os campos!");
    return;
  }

  // Cria o objeto da transação
  const transaction = { description, category,  amount, type, payment, obs, date };

  // Envia a transação para o backend
  addTransaction(transaction);

  // Limpa o formulário
  form.reset();
});

// Carrega as transações ao abrir a página
fetchTransactions();
