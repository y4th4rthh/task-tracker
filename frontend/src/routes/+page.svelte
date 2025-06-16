<script>
  import "../app.css";
  let email = '';
  let password = '';

  async function handleLogin() {
    const res = await fetch('http://localhost:8080/api/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password }),
    });

    const data = await res.json();
    console.log(data);
  }
</script>

<div class="max-w-md mx-auto mt-20 p-4 bg-black text-gray-200 rounded shadow">
  <h2 class="text-xl font-bold mb-4">Login</h2>
  <input class="w-full mb-2 p-2 border" bind:value={email} placeholder="Email" />
  <input class="w-full mb-4 p-2 border" type="password" bind:value={password} placeholder="Password" />
  <button on:click={handleLogin} class="bg-blue-500 text-white px-4 py-2 rounded">Login</button>
</div>
