export type Task = {
  id: number
  title: string
  created_at: Date
  updated_at: Date
}

export type CsrfToke = {
  csrf_token: string
}

export type Credential = {
  email: string
  password: string
}
