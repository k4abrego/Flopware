const API_URL =
  import.meta.env.VITE_API_URL ?? "http://localhost:8080/api";

export async function getProjects() {
  const response = await fetch(`${API_URL}/projects`);

  if (!response.ok) {
    throw new Error("No fue posible obtener los proyectos");
  }

  return response.json();
}