<script lang="ts">
  import { enhance } from "$app/forms";
  export let data;
</script>

<div class="flex">
  {#each data.domains as domain}
    <div class="p-4 border rounded m-2">
      <a href={`/panel/${domain.id}`}>
        <h2 class="text-xl font-bold">{domain.name}</h2>
        <p>{domain.networks.length} Networks</p>
      </a>
      <form method="POST" action="?/deleteDomain" class="mt-2" use:enhance>
        <input type="hidden" name="domainId" value={domain.id} />
        <button
          type="submit"
          class="bg-red-500 hover:bg-red-600 text-white font-semibold py-1 px-3 rounded shadow transition-colors duration-150 cursor-pointer"
          >Delete</button
        >
      </form>
    </div>
  {/each}
  <br />
  <div class="p-4 border rounded m-2">
    <h2 class="text-xl font-bold">Add New Domain</h2>
    <form method="POST" action="?/createDomain" use:enhance>
      <input
        name="domainName"
        type="text"
        placeholder="Domain Name"
        class="border p-2 rounded w-full mb-2"
      />
      <button type="submit" class="bg-blue-500 text-white p-2 rounded"
        >Add Domain</button
      >
    </form>
  </div>
</div>
