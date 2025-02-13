// URL da API do backend
const API_URL = "http://localhost:8080";

// Referências aos elementos HTML
const form = document.getElementById("category-form");
const categoriesTable = document.getElementById("category-table");

// Função para exibir transações na tabela
function displayCategories(categories) {
  categoriesTable.innerHTML = ""; // Limpa a tabela

  categories.forEach((category) => {
    const row = document.createElement("tr");

    row.innerHTML = `
      <td>${category.category}</td>
    `;

    categoriesTable.appendChild(row);
  });
}

async function fetchCategories() {
  try {
    const response = await fetch(`${API_URL}/categorias`);
    if (!response.ok) throw new Error("Erro ao buscar categorias.");

    const categories = await response.json();

    if (Array.isArray(categories)) {
        categories.forEach(category => {
            if (category.category !== undefined) {
                console.log(`Categoria: ${category.category}`);
                displayCategories(categories);
            } else {
                console.warn("Categoria null:", category);
            }
        });
    } else {
        console.error("Formato inesperado de resposta da API:", categories);
    }
    
  } catch (error) {
    console.error(error.message);
  }
}

// Lida com o envio do formulário
form.addEventListener("submit", (event) => {
  event.preventDefault();

  // Obtém os valores do formulário
  const category = document.getElementById("category").value;

  if (!category) {
    alert("Preencha todos os campos!");
    return;
  }

  const category_object = {
    category,
  };

  // Envia a transação para o backend
  addCategory(category_object);

  // Limpa o formulário
  form.reset();
});

// Função para adicionar uma nova transação
async function addCategory(category) {
  try {
    const response = await fetch(`${API_URL}/categorias`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(category),
    });

    if (!response.ok) throw new Error("Erro ao adicionar transação.");
    fetchCategories(); // Atualiza a lista após adicionar
  } catch (error) {
    console.error(error.message);
  }
}

fetchCategories(); // Busca as categorias ao carregar a página