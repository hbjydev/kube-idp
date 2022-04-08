<template>
  <div class="navbar-wrapper">

    <div class="grid grid-cols-3 w-full px-4 gap-8">
      <div class="flex items-center font-bold">
        kube-idp
      </div>
      <div class="flex items-center">
        <input
          type="text"
          id="searchbox"
          class="h-9 w-full bg-gray-700 text-gray-200 rounded px-4 outline-none focus:h-12 transition-all duration-200"
          placeholder="Search..."
        />
      </div>
      <div class="flex items-center justify-end gap-2">
        <div class="flex items center justify-end gap-2" v-if="!loading">
          <div class="flex flex-col items-end">
            <span>Welcome back, <b>{{ result.user.displayName }}</b>!</span>
            <a class="text-sm" href="#">Not you? <b>Log out</b>.</a>
          </div>
          <img class="h-10 rounded-full" src="https://picsum.photos/seed/picsum/128" alt="Profile icon">
        </div>
      </div>
    </div>

  </div>
</template>

<script lang="ts">
import { useQuery } from '@vue/apollo-composable';
import gql from 'graphql-tag';

interface User {
  id: string;
  login: string;
  displayName: string;
  email: string;
  createdAt: Date;
  updatedAt: Date;
}

interface GetUserByLoginData {
  user: User;
}
interface GetUserByLoginVars {
  login: string;
}

const GET_USER_BY_LOGIN = gql`
  query getUserByLogin($login: String!) {
    user(login: $login) {
      id
      login
      displayName
      email
      createdAt
      updatedAt
    }
  }
`;


export default {
  setup() {
    const { result, loading, error } = useQuery<GetUserByLoginData, GetUserByLoginVars>(GET_USER_BY_LOGIN, { login: 'admin' });
    return { result, loading, error };
  }
}

</script>

<style scoped lang="postcss">
  .navbar-wrapper {
    @apply bg-gray-900 text-white h-16 flex items-center;
  }
</style>
