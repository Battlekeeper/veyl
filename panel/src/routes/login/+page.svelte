<script lang="ts">
import { enhance } from "$app/forms";
    export let data;    
</script>

<h1 class="text-2xl font-bold mb-4">Login</h1>
<form method="POST" action="?/login" use:enhance>
    <input name="email" type="email" placeholder="Email" class="border p-2 rounded w-full mb-2" required />
    <input name="password" type="password" placeholder="Password" class="border p-2 rounded w-full mb-2" required />
    <button type="submit" class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded shadow transition-colors duration-150 cursor-pointer">Login</button>
</form>