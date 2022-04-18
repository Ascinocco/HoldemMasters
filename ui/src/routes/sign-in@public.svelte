<script lang="ts">
  type FormErrors = {
    email?: string;
    password?: string;
  } | undefined;

  let email = "";
  let password = "";
  let errors;

  const validateFormData = () => {
    let errors: FormErrors = undefined;

    if (!email) {
      errors = {
        email: "Email is required"
      }
    }

    if (!password) {
      errors = {
        ...errors,
        password: "Password is required",
      }
    }

    return errors;
  };

  const handleSubmit = (e) => {
    errors = validateFormData();
    if (errors) return;
  };
</script>

<div class="flex justify-center h-screen items-center">
  <div class="card w-96 bg-base-100 bg-slate-400 shadow-xl h-1/3 text-black nav-offset card-min-height">
    <div class="card-body">
      <div class="justify-between">
        <h1 class="card-title pb-5">Sign In</h1>
        <form id="sign-in" on:submit|preventDefault={handleSubmit}>
          <div class="form-control w-full max-w-xs">
            <label for="email" class="label">
              <span class="label-text-alt text-black">Email</span>
            </label>
            <input bind:value={email} name="email" type="email" placeholder="Email" class="input input-bordered w-full max-w-xs text-slate-400">
            <span class="label-text-alt text-rose-800 p-1 input-error">
              {#if errors && errors.email}
                {errors.email}
              {/if}
            </span>
          </div>
          <div class="form-control w-full max-w-xs pb-14">
            <label for="password" class="label">
              <span class="label-text-alt text-black">Password</span>
            </label>
            <input bind:value={password} name="password" type="password" placeholder="Password" class="input input-bordered w-full max-w-xs text-slate-400">
            <span class="label-text-alt text-rose-800 p-1 input-error">
              {#if errors && errors.password}
                {errors.password}
              {/if}
            </span>
          </div>
          <div class="card-actions">
            <button for="sign-in" class="btn btn-primary btn-block" type="submit">Sign In</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</div>

<style>
  .nav-offset {
    margin-bottom: 30vh;
  }

  .input-error {
    min-height: 16px;
  }

  .card-min-height {
    min-height: 400px;
  }
</style>