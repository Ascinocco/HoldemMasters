<script lang="ts">
  import { goto } from "$app/navigation";
  import { user } from "../stores/user";
  type FormErrors = {
    email?: string;
    username?: string;
    password?: string;
    passwordConfirmation?: string;
  } | undefined;

  type CreateUserResponse = {
    email: string;
    username: string;
    token: string;
    error: string;
  };

  const pwRegex = /^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$/;

  let email = "";
  let username = "";
  let password = "";
  let passwordConfirmation = "";
  let errors;
  let apiError;

  const validateFormData = () => {
    let errors: FormErrors = undefined;

    if (!email) {
      errors = {
        email: "Email is required."
      }
    }

    if (!username) {
      errors = {
        ...errors,
        username: "Username is required."
      }
    }

    if (!password) {
      errors = {
        ...errors,
        password: "Password is required.",
      }
    }

    if (!pwRegex.test(password)) {
      errors = {
        ...errors,
        password: "Password must be at least 8 characters, and contain a number and symbol.",
      }
    }

    if (!passwordConfirmation) {
      errors = {
        ...errors,
        passwordConfirmation: "Please re-enter your password."
      }
    }

    if (password !== passwordConfirmation) {
      errors = {
        ...errors,
        passwordConfirmation: "Passwords do not match."
      }
    }

    return errors;
  };

  const createUser = () => {
    fetch("http://localhost:8080/api/v1/user", {
      method: "POST",
      body: JSON.stringify({
        email,
        username,
        password,
        passwordConfirmation
      })
    })
    .then(res => res.json())
    .then((res: CreateUserResponse) => {
      if (res.error) {
        apiError = res.error;
        return;
      }

      apiError = null;
      sessionStorage.setItem("token", res.token);
      user.add({
        email: res.email,
        username: res.username,
      });

      goto("/app/play");
    })
  }

  const handleSubmit = () => {
    errors = validateFormData();
    if (errors) return;
    createUser();
  };
</script>

<div class="flex justify-center h-screen items-center">
  <div class="card w-96 bg-base-100 bg-slate-400 shadow-xl text-black nav-offset card-min-height">
    <div class="card-body">
      <div class="justify-between">
        <h1 class="card-title pb-5">Sign Up</h1>
        <form id="sign-up" on:submit|preventDefault={handleSubmit}>
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
          <div class="form-control w-full max-w-xs">
            <label for="Username" class="label">
              <span class="label-text-alt text-black">Username</span>
            </label>
            <input bind:value={username} name="username" type="text" placeholder="Username" class="input input-bordered w-full max-w-xs text-slate-400">
            <span class="label-text-alt text-rose-800 p-1 input-error">
              {#if errors && errors.username}
                {errors.username}
              {/if}
            </span>
          </div>
          <div class="form-control w-full max-w-xs">
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
          <div class="form-control w-full max-w-xs pb-8">
            <label for="password-confirmation" class="label">
              <span class="label-text-alt text-black">Password Confirmation</span>
            </label>
            <input bind:value={passwordConfirmation} name="password-confirmation" type="password" placeholder="Password Confirmation" class="input input-bordered w-full max-w-xs text-slate-400">
            <span class="label-text-alt text-rose-800 p-1 input-error">
              {#if errors && errors.passwordConfirmation}
                {errors.passwordConfirmation}
              {/if}
            </span>
          </div>
          <div class="card-actions">
            <span class="label-text-alt text-rose-800 p-1 input-error">
              {#if apiError}
                {apiError}
              {/if}
            </span>
            <button for="sign-up" class="btn btn-primary btn-block" type="submit">Sign Up</button>
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