<!-- eslint-disable no-extra-semi -->
<!-- eslint-disable no-extra-semi -->
<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-5">Login</h1>
        <br />
        <form-tag @myevent="submitHandler" name="myform" event="myevent">
          <text-input
            label="Email"
            type="email"
            name="email"
            required="true"
            v-model="email"
          />
          <text-input
            label="Password"
            type="password"
            name="password"
            required="true"
            v-model="password"
          />
          <input type="submit" class="btn btn-primary" value="Login" />
        </form-tag>
      </div>
    </div>
  </div>
</template>

<script>
import TextInput from "../components/Forms/TextInput.vue"
import FormTag from "../components/Forms/FormTag.vue"

export default {
  name: "Login",
  components: {
    "text-input": TextInput,
    "form-tag": FormTag,
  },
  data() {
    return {
      email: "",
      password: "",
    }
  },
  methods: {
    submitHandler() {
      console.log(
        "submitHandler called - success" + this.email + " " + this.password
      )
    },
  },
}
</script>

<style></style>
