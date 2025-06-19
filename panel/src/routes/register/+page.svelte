<script lang="ts">
  import { enhance } from "$app/forms";
  export let data;
</script>

<div class="w-full h-screen flex flex-col gap-4 justify-center items-center">
  <form method="POST" class="" use:enhance>
    <label>
      Email
      <input name="email" type="email" class="border rounded-sm" required />
    </label>
    <label>
      Password
      <input
        name="password"
        type="password"
        class="border rounded-sm"
        required
      />
    </label>
    <label>
      Confirm Password
      <input
        name="confirm_password"
        type="password"
        class="border rounded-sm"
        required
      />
    </label>
    <button>Signup</button>
  </form>
</div>
