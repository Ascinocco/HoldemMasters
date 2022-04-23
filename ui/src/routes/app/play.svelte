<script lang="ts">
  import { user } from "../../stores/user";
  const email = $user.email;
  const username = $user.username; 
</script>

<div>
  <h2>{email || "not set"}</h2>
  <h2>{username || "not set"}</h2>
  Logged in
</div>