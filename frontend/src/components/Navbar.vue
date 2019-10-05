<template>
  <nav class="navbar navbar-dark bg-frontend">
    <a class="navbar-brand" href="/">frontend</a>

    <span class="navbar-text">
      {{ currentUser.email }}
      <router-link to="/logout" class="btn btn-logout btn-sm">Logout</router-link>
    </span>

  </nav>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'Navbar',
  computed: {
    ...mapGetters({ currentUser: 'currentUser' })
  }
}
</script>

<style scoped>
.bg-frontend {
  background: #007bff;
}

.btn-logout {
  background-color: #333333;
  color: #fff;
}

.btn-logout:hover {
  background-color: darken(#837A75, 20%);
  border-color: darken(#837A75, 20%);
  color: #fff;
}

.navbar-text {
  color: #fff;
}
</style>
