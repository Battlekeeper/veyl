<script lang="ts">
  import { enhance } from "$app/forms";
  export let data;
</script>

<!-- <pre>{JSON.stringify(data, null, 4)}</pre> -->
<h1 class="text-2xl font-bold mb-4">{data.domain.name}</h1>
<h1>Networks</h1>
<div class="flex">
  {#each data.networks as network}
    <div class="p-4 border rounded m-2">
      <a href={`/panel/${data.domain.id}/${network.id}`}>
        <h2 class="text-xl font-bold">{network.name}</h2>
      </a>
      <p>{network.resources.length} Resources</p>
      <p>{network.relays.length} Relays</p>
      <form method="POST" action="?/deleteNetwork" class="mt-2" use:enhance>
        <input type="hidden" name="networkId" value={network.id} />
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
    <h2 class="text-xl font-bold">Add New Network</h2>
    <form method="POST" action="?/createNetwork" use:enhance>
      <input
        name="networkName"
        type="text"
        placeholder="Network Name"
        class="border p-2 rounded w-full mb-2"
      />
      <button
        type="submit"
        class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded shadow transition-colors duration-150 cursor-pointer"
        >Add Network</button
      >
    </form>
  </div>
</div>
