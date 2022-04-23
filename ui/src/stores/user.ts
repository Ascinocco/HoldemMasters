import { writable } from 'svelte/store';

type User = {
	email?: string;
	username?: string;
};

const defaultUser: User = {
	email: '',
	username: ''
};

const createUser = () => {
	const { subscribe, set, update } = writable<User>(defaultUser);

	return {
		subscribe,
		add: (user: User) =>
			update((ou) => ({
				...ou,
				...user
			})),
		clear: () => set(defaultUser)
	};
};

export const user = createUser();
