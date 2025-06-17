<script lang="ts">
  export let data;
</script>

<div class="flex">
  {#each data.domains as domain}
    <div class="p-4 border rounded m-2">
      <h2 class="text-xl font-bold">{domain.name}</h2>
      <p>{domain.networks.length} Networks</p>
    </div>
  {/each}
  <br>
  <div class="p-4 border rounded m-2">
    <h2 class="text-xl font-bold">Add New Domain</h2>
    <form method="POST" action="?/createDomain">
      <input name="domainName" type="text" placeholder="Domain Name" class="border p-2 rounded w-full mb-2" />
      <button type="submit" class="bg-blue-500 text-white p-2 rounded">Add Domain</button>
    </form>
  </div>
</div>
