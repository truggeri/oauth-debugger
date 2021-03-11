<script lang="typescript">
  export let user: { username: string, password: string}

  function authorize() { alert("Authorize " + user.username)}
  function deny() {}
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
    <button class="authorize" on:click|once={authorize}>Authorize</button>
    <button on:click|once={deny}>Deny</button>
  </div>
</div>