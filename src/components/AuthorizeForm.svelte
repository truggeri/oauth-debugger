<script lang="typescript">
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();

  export let user: { username: string, password: string}
</script>

<style>
  .authorize {
    background: linear-gradient(green, #009900);
    background-color: green;
  }

  .authorize:hover {
    background: linear-gradient(#009900, #00bb00);
  }
</style>

<div>
  <p><em>OAuth Debugger</em> would like permissions to access your account, {user.username}</p>

  <div>
    <h3>Review Permissions</h3>
    <ul>
      <li>User Information</li>
      <li>Profile Picture</li>
    </ul>
  </div>
  <br />

  <div>
    <button on:click|once={() => dispatch("submit", {})} class="authorize">Authorize</button>
    <button on:click|once={() => dispatch("deny", {})}>Deny</button>
  </div>
</div>