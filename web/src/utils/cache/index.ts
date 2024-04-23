class Cache {
  set(key: string, value: any) {
    window.localStorage.setItem(key, JSON.stringify(value));
  }

  get(key: string) {
    const result = window.localStorage.getItem(key);
    if (result) {
      return JSON.parse(result);
    }
  }

  remove(key: string) {
    return window.localStorage.removeItem(key);
  }

  clear() {
    window.localStorage.clear();
  }
}

export const cache = new Cache();
