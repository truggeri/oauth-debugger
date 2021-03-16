<script lang="typescript">
  import Alert from "./Alert.svelte";
  export let client: any;
</script>

<div>
  <Alert klass="success" boldMsg="Success" message="Here are your credentials" />
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