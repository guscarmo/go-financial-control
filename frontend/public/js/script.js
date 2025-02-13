// URL da API do backend
const API_URL = "http://localhost:8080";
const categorySelect = document.getElementById("category"); // Dropdown de categorias

// Referências aos elementos HTML
const form = document.getElementById("transaction-form");
const transactionsTable = document.getElementById("transactions-table");

// Função para buscar e exibir transações
async function fetchTransactions() {
  try {
    const response = await fetch(`${API_URL}/transacoes`);
    if (!response.ok) throw new Error("Erro ao buscar transações.");

    const transactions = await response.json();

    if (Array.isArray(transactions)) {
        transactions.forEach(transaction => {
            if (transaction.amount !== undefined) {
                console.log(`Descrição: ${transaction.description}`);
                console.log(`Valor: ${transaction.amount.toFixed(2)}`);
                console.log(`Date: ${transaction.date}`); // Aqui usamos toFixed corretamente
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
      <td>${new Date(transaction.date).toLocaleDateString()}</td>
    `;

    transactionsTable.appendChild(row);
  });
}

// Função para adicionar uma nova transação
async function addTransaction(transaction) {
  try {
    const response = await fetch(`${API_URL}/transacoes`, {
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
  const typ = document.getElementById("typ").value;
  const payment = document.getElementById("payment").value;
  const obs = document.getElementById("obs").value;


  if (!description || !amount || !date || !category || !typ || !payment) {
    alert("Preencha todos os campos!");
    return;
  }

  // Cria o objeto da transação
  const transaction = { description, category,  amount, typ, payment, obs, date };

  // Envia a transação para o backend
  addTransaction(transaction);

  // Limpa o formulário
  form.reset();
});

// Função para buscar categorias e preencher o dropdown
async function fetchCategories() {
  try {
    const response = await fetch(`${API_URL}/categorias`);
    if (!response.ok) throw new Error("Erro ao buscar categorias.");

    const categories = await response.json();

    // Preenche o dropdown com as categorias
    categorySelect.innerHTML = ""; // Limpa o dropdown
    categories.forEach((category) => {
      const option = document.createElement("option");
      option.value = category.id;
      option.textContent = category.category;
      categorySelect.appendChild(option);
    });
  } catch (error) {
    console.error(error.message);
  }
}

// Carrega as transações ao abrir a página
fetchTransactions();
fetchCategories();
