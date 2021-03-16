<script lang="typescript">
  import { blur } from "svelte/transition";

  export let client: any;
</script>

<style>
  div.success {
    background-color: rgba(0, 255, 0, 0.3);
    border: 2px solid #009900;
    border-radius: 6px;
    margin-bottom: 12px;
    padding: 12px;
    width: 100%;
  }
</style>

<div>
  <div class="success" transition:blur="{{duration: 500}}">
    <strong>Success!</strong> Here are your credentials.
  </div>

  <br>

  <table>
    <thead>
      <tr>
        <th>Field</th>
        <th>Value</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>App Name</td>
        <td>{ client.name }</td>
      </tr>
      <tr>
        <td>Redirect Uri</td>
        <td>{ client.redirect_uri }</td>
      </tr>
      <tr>
        <td>Client Id</td>
        <td>{ client.data.ClientId }</td>
      </tr>
      <tr>
        <td>Client Secret</td>
        <td>{ client.data.ClientSecret }</td>
      </tr>
      <tr>
        <td>Expires</td>
        <td>{ client.data.Expires }</td>
      </tr>
    </tbody>
  </table>
</div>