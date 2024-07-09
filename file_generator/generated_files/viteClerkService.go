// THIS FILE IS AUTOGENERATED, DO NOT MODIFY

package generated

var File__viteClerkService = []string{
"import { User } from '@clerk/clerk-js';",
"import Clerk from '@clerk/clerk-js';",
"import axios from 'axios';",
"",
"const clerk = new Clerk(import.meta.env.VITE_CLERK_PUBLISHABLE_KEY);",
"clerk.load();",
"",
"const API_URL = import.meta.env.VITE_API_URL;",
"",
"export const userService = {",
"  getCurrentUser: async (): Promise<User | null> => {",
"    try {",
"      const clerkUser = await clerk.user.getCurrent();",
"      if (clerkUser) {",
"        const token = await clerk.session.getToken();",
"        const response = await axios.get(`${API_URL}/users/current`, {",
"          headers: { Authorization: `Bearer ${token}` }",
"        });",
"        return response.data;",
"      }",
"      return null;",
"    } catch (error) {",
"      console.error('Error getting current user:', error);",
"      return null;",
"    }",
"  },",
"",
"  signIn: async (): Promise<void> => {",
"    try {",
"      await clerk.openSignIn();",
"      const user = await clerk.user.getCurrent();",
"      if (user) {",
"        const token = await clerk.session.getToken();",
"        await axios.post(`${API_URL}/users`, user, {",
"          headers: { Authorization: `Bearer ${token}` }",
"        });",
"      }",
"    } catch (error) {",
"      console.error('Error during sign in:', error);",
"    }",
"  },",
"",
"  signOut: async (): Promise<void> => {",
"    try {",
"      await clerk.signOut();",
"    } catch (error) {",
"      console.error('Error during sign out:', error);",
"    }",
"  },",
"",
"  getUserToken: async (): Promise<string | null> => {",
"    try {",
"      return await clerk.session.getToken();",
"    } catch (error) {",
"      console.error('Error getting user token:', error);",
"      return null;",
"    }",
"  },",
"",
"  isAuthenticated: (): boolean => {",
"    return !!clerk.user;",
"  },",
"",
"  updateUser: async (userData: Partial<User>): Promise<User | null> => {",
"    try {",
"      const token = await clerk.session.getToken();",
"      const response = await axios.put(`${API_URL}/users`, userData, {",
"        headers: { Authorization: `Bearer ${token}` }",
"      });",
"      return response.data;",
"    } catch (error) {",
"      console.error('Error updating user:', error);",
"      return null;",
"    }",
"  }",
"};",
}