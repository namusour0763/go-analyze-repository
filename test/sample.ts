interface User {
  id: number;
  name: string;
  email: string;
}

class UserService {
  private users: User[] = [];

  addUser(user: User): void {
    this.users.push(user);
  }

  getUser(id: number): User | undefined {
    return this.users.find(user => user.id === id);
  }

  getAllUsers(): User[] {
    return this.users;
  }
}

const userService = new UserService();
userService.addUser({ id: 1, name: "John", email: "john@example.com" });
userService.addUser({ id: 2, name: "Jane", email: "jane@example.com" });

console.log(userService.getAllUsers());