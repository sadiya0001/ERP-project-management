<template>
  <aside class="sidebar">
    <h2>ERP System</h2>

    <ul>
      <li>Dashboard</li>
      <li>Projects</li>
      <li>Teams</li>
      <li>Tasks</li>
    </ul>
  </aside>
</template>

<style scoped>
.sidebar {
  width: 240px;
  min-height: 100vh;
  background: #1e293b;
  color: white;
  padding: 20px;
}

.sidebar h2 {
    margin-bottom: 30px;
}

.sidebar ul {
  list-style: none;
}

.sidebar li {
  margin-bottom: 15px;
  cursor: pointer;
}
</style>
