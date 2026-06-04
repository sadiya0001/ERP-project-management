import { defineStore } from 'pinia'
import api from '../services/api'

export const useProjectStore = defineStore('project', {
  state: () => ({
    projects: [],
    currentProject: null,
    tasks: [],
    teamMembers: [],
    loading: false,
    error: null
  }),

  getters: {
    projectCount: (state) => state.projects.length,
    tasksByStatus: (state) => {
      const counts = { completed: 0, onHold: 0, inProgress: 0, pending: 0 }
      state.tasks.forEach(task => {
        if (task.status === 'Completed') counts.completed++
        else if (task.status === 'On-hold') counts.onHold++
        else if (task.status === 'In Progress') counts.inProgress++
        else if (task.status === 'Pending') counts.pending++
      })
      return counts
    }
  },

  actions: {
    async fetchProjects() {
      this.loading = true
      try {
        const response = await api.get('/projects')
        this.projects = response.data
      } catch (err) {
        // Load mock data
        this.projects = getMockProjects()
      } finally {
        this.loading = false
      }
    },

    async fetchProject(id) {
      this.loading = true
      try {
        const response = await api.get(`/projects/${id}`)
        this.currentProject = response.data
      } catch (err) {
        this.currentProject = getMockProjects().find(p => p.id == id) || getMockProjects()[0]
      } finally {
        this.loading = false
      }
    },

    async createProject(projectData) {
      this.loading = true
      try {
        const response = await api.post('/projects', projectData)
        this.projects.push(response.data)
        return response.data
      } catch (err) {
        const newProject = {
          id: Date.now(),
          ...projectData,
          status: 'OnTrack',
          members: [],
          issues: 0,
          createdAt: new Date().toISOString()
        }
        this.projects.push(newProject)
        return newProject
      } finally {
        this.loading = false
      }
    },

    async fetchTasks(projectId) {
      this.loading = true
      try {
        const response = await api.get(`/projects/${projectId}/tasks`)
        this.tasks = response.data
      } catch (err) {
        this.tasks = getMockTasks()
      } finally {
        this.loading = false
      }
    },

    async createTask(projectId, taskData) {
      this.loading = true
      try {
        const response = await api.post(`/projects/${projectId}/tasks`, taskData)
        this.tasks.push(response.data)
        return response.data
      } catch (err) {
        const newTask = {
          id: Date.now(),
          taskId: `#NV${Math.floor(Math.random() * 999)}`,
          ...taskData,
          status: 'Pending',
          createdAt: new Date().toISOString(),
          createdBy: 'Yash Ghori',
          assignee: { name: 'Yash Ghori', avatar: 'https://i.pravatar.cc/150?img=11' }
        }
        this.tasks.push(newTask)
        return newTask
      } finally {
        this.loading = false
      }
    },

    async fetchTeamMembers() {
      this.loading = true
      try {
        const response = await api.get('/team')
        this.teamMembers = response.data
      } catch (err) {
        this.teamMembers = getMockTeamMembers()
      } finally {
        this.loading = false
      }
    }
  }
})

function getMockProjects() {
  return [
    {
      id: 1,
      title: 'Adoddle',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam.',
      status: 'Offline',
      date: '05 APRIL 2023',
      issues: 14,
      color: '#4a7cff',
      members: [
        { id: 1, name: 'Yash Ghori', avatar: 'https://i.pravatar.cc/150?img=11' },
        { id: 2, name: 'Priya Shah', avatar: 'https://i.pravatar.cc/150?img=5' },
        { id: 3, name: 'Raj Patel', avatar: 'https://i.pravatar.cc/150?img=12' },
        { id: 4, name: 'Anita Kumar', avatar: 'https://i.pravatar.cc/150?img=9' }
      ]
    },
    {
      id: 2,
      title: 'Adoddle',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam.',
      status: 'OnTrack',
      date: '05 APRIL 2023',
      issues: 14,
      color: '#22c55e',
      members: [
        { id: 5, name: 'Sonia Verma', avatar: 'https://i.pravatar.cc/150?img=20' },
        { id: 6, name: 'Amit Sharma', avatar: 'https://i.pravatar.cc/150?img=14' },
        { id: 7, name: 'Neha Gupta', avatar: 'https://i.pravatar.cc/150?img=23' }
      ]
    },
    {
      id: 3,
      title: 'Adoddle',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam.',
      status: 'Offline',
      date: '05 APRIL 2023',
      issues: 14,
      color: '#f97316',
      members: [
        { id: 8, name: 'Vikram Singh', avatar: 'https://i.pravatar.cc/150?img=33' },
        { id: 9, name: 'Divya Joshi', avatar: 'https://i.pravatar.cc/150?img=25' },
        { id: 10, name: 'Karan Mehta', avatar: 'https://i.pravatar.cc/150?img=51' },
        { id: 11, name: 'Pooja Reddy', avatar: 'https://i.pravatar.cc/150?img=44' }
      ]
    },
    {
      id: 4,
      title: 'Adoddle',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam.',
      status: 'Offline',
      date: '05 APRIL 2023',
      issues: 14,
      color: '#ef4444',
      members: [
        { id: 12, name: 'Arjun Nair', avatar: 'https://i.pravatar.cc/150?img=53' },
        { id: 13, name: 'Meera Das', avatar: 'https://i.pravatar.cc/150?img=47' },
        { id: 14, name: 'Rahul Iyer', avatar: 'https://i.pravatar.cc/150?img=57' }
      ]
    },
    {
      id: 5,
      title: 'Adoddle',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam.',
      status: 'OnTrack',
      date: '05 APRIL 2023',
      issues: 14,
      color: '#8b5cf6',
      members: [
        { id: 15, name: 'Sneha Rao', avatar: 'https://i.pravatar.cc/150?img=38' },
        { id: 16, name: 'Manish Tiwari', avatar: 'https://i.pravatar.cc/150?img=60' }
      ]
    },
    {
      id: 6,
      title: 'Adoddle',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam.',
      status: 'Offline',
      date: '05 APRIL 2023',
      issues: 14,
      color: '#06b6d4',
      members: [
        { id: 17, name: 'Ritu Kapoor', avatar: 'https://i.pravatar.cc/150?img=41' },
        { id: 18, name: 'Deepak Pandey', avatar: 'https://i.pravatar.cc/150?img=52' },
        { id: 19, name: 'Kavya Bhatt', avatar: 'https://i.pravatar.cc/150?img=43' },
        { id: 20, name: 'Sunil Mishra', avatar: 'https://i.pravatar.cc/150?img=59' }
      ]
    }
  ]
}

function getMockTasks() {
  return [
    {
      id: 1,
      taskId: '#NV221',
      title: 'Make an Automatic Payment System that enable the design',
      status: 'Completed',
      timeSpent: '2h 30m',
      deadline: '2023-04-20',
      createdAt: '2023-04-06',
      daysAgo: 14,
      createdBy: 'Yash Ghori',
      assignee: { name: 'Yash Ghori', avatar: 'https://i.pravatar.cc/150?img=11' }
    },
    {
      id: 2,
      taskId: '#NV222',
      title: 'Make an Automatic Payment System that enable the design',
      status: 'On-hold',
      timeSpent: '1h 15m',
      deadline: '2023-04-22',
      createdAt: '2023-04-06',
      daysAgo: 14,
      createdBy: 'Yash Ghori',
      assignee: { name: 'Priya Shah', avatar: 'https://i.pravatar.cc/150?img=5' }
    },
    {
      id: 3,
      taskId: '#NV223',
      title: 'Make an Automatic Payment System that enable the design',
      status: 'Completed',
      timeSpent: '4h 00m',
      deadline: '2023-04-25',
      createdAt: '2023-04-06',
      daysAgo: 14,
      createdBy: 'Yash Ghori',
      assignee: { name: 'Raj Patel', avatar: 'https://i.pravatar.cc/150?img=12' }
    },
    {
      id: 4,
      taskId: '#NV224',
      title: 'Make an Automatic Payment System that enable the design',
      status: 'In Progress',
      timeSpent: '0h 45m',
      deadline: '2023-04-28',
      createdAt: '2023-04-06',
      daysAgo: 14,
      createdBy: 'Yash Ghori',
      assignee: { name: 'Anita Kumar', avatar: 'https://i.pravatar.cc/150?img=9' }
    },
    {
      id: 5,
      taskId: '#NV225',
      title: 'Make an Automatic Payment System that enable the design',
      status: 'Pending',
      timeSpent: '0h 00m',
      deadline: '2023-04-30',
      createdAt: '2023-04-06',
      daysAgo: 14,
      createdBy: 'Yash Ghori',
      assignee: { name: 'Sonia Verma', avatar: 'https://i.pravatar.cc/150?img=20' }
    },
    {
      id: 6,
      taskId: '#NV226',
      title: 'Make an Automatic Payment System that enable the design',
      status: 'Completed',
      timeSpent: '3h 20m',
      deadline: '2023-05-02',
      createdAt: '2023-04-06',
      daysAgo: 14,
      createdBy: 'Yash Ghori',
      assignee: { name: 'Amit Sharma', avatar: 'https://i.pravatar.cc/150?img=14' }
    },
    {
      id: 7,
      taskId: '#NV227',
      title: 'Make an Automatic Payment System that enable the design',
      status: 'On-hold',
      timeSpent: '1h 50m',
      deadline: '2023-05-05',
      createdAt: '2023-04-06',
      daysAgo: 14,
      createdBy: 'Yash Ghori',
      assignee: { name: 'Neha Gupta', avatar: 'https://i.pravatar.cc/150?img=23' }
    }
  ]
}

function getMockTeamMembers() {
  const names = [
    'Yash Ghori', 'Priya Shah', 'Raj Patel', 'Anita Kumar', 'Sonia Verma',
    'Amit Sharma', 'Neha Gupta', 'Vikram Singh', 'Divya Joshi', 'Karan Mehta',
    'Pooja Reddy', 'Arjun Nair', 'Meera Das', 'Rahul Iyer', 'Sneha Rao',
    'Manish Tiwari', 'Ritu Kapoor', 'Deepak Pandey', 'Kavya Bhatt', 'Sunil Mishra',
    'Nandini Bose', 'Aakash Gill', 'Tanvi Desai', 'Rohan Malhotra', 'Ishita Sen',
    'Varun Chopra', 'Simran Kaur', 'Nikhil Saxena', 'Aditi Thakur', 'Gaurav Sinha',
    'Pallavi Kulkarni', 'Siddharth Jain', 'Roshni Dutta', 'Abhishek Banerjee', 'Lakshmi Pillai',
    'Harsh Agarwal', 'Swati Chauhan', 'Pranav Goyal', 'Megha Bhat'
  ]
  const avatarIds = [11, 5, 12, 9, 20, 14, 23, 33, 25, 51, 44, 53, 47, 57, 38, 60, 41, 52, 43, 59, 1, 2, 3, 4, 6, 7, 8, 10, 13, 15, 16, 17, 18, 19, 21, 22, 24, 26, 27]
  const ringColors = ['#4a7cff', '#22c55e', '#ef4444', '#f97316', '#8b5cf6', '#06b6d4', '#ec4899', '#eab308']

  return names.map((name, i) => ({
    id: i + 1,
    name,
    avatar: `https://i.pravatar.cc/150?img=${avatarIds[i]}`,
    role: i === 0 ? 'Team Lead' : 'Developer',
    ringColor: ringColors[i % ringColors.length]
  }))
}
