import { type FormEvent, useState } from "react"

import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"

type DownloadResponse = {
  links: string[]
}

function App() {
  const [url, setUrl] = useState("")
  const [links, setLinks] = useState<string[] | null>(null)
  const [error, setError] = useState<string | null>(null)
  const [isLoading, setIsLoading] = useState(false)

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault()
    if (!url) {
      setError("Introduce la URL de la serie en animeav1.com.")
      setLinks(null)
      return
    }

    setIsLoading(true)
    setError(null)
    setLinks(null)

    try {
      const response = await fetch("/api/download-links", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ url }),
      })

      if (!response.ok) {
        let message = "No se pudo obtener la lista de enlaces."
        try {
          const data = await response.json()
          if (typeof data?.error === "string") {
            message = data.error
          }
        } catch (parseError) {
          // el endpoint puede no devolver body en caso de error
        }
        throw new Error(message)
      }

      const data = (await response.json()) as DownloadResponse
      setLinks(data.links ?? [])
    } catch (fetchError) {
      const message =
        fetchError instanceof Error
          ? fetchError.message
          : "Se produjo un error desconocido."
      setError(message)
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <main className="min-h-screen bg-background">
      <div className="mx-auto flex min-h-screen w-full max-w-xl flex-col items-center justify-center px-4 py-8">
        <Card className="w-full">
          <CardHeader>
            <CardTitle>AnimeAV1 Downloader</CardTitle>
            <CardDescription>
              Introduce la URL de la serie en animeav1.com para obtener los enlaces de descarga desde Mega.
            </CardDescription>
          </CardHeader>

          <CardContent>
            <form className="space-y-4" onSubmit={handleSubmit}>
              <div className="space-y-2">
                <Label htmlFor="url">URL de la serie</Label>
                <Input
                  id="url"
                  type="url"
                  placeholder="https://animeav1.com/media/..."
                  value={url}
                  onChange={(event) => setUrl(event.target.value)}
                  required
                />
              </div>

              <Button type="submit" disabled={isLoading} className="w-full">
                {isLoading ? "Buscando enlaces..." : "Obtener enlaces"}
              </Button>
            </form>

            {error && (
              <p className="mt-4 rounded-md border border-destructive/40 bg-destructive/10 px-3 py-2 text-sm text-destructive">
                {error}
              </p>
            )}

            {links && (
              <div className="mt-6 space-y-2">
                <h2 className="text-base font-semibold">Enlaces de Mega</h2>
                {links.length === 0 ? (
                  <p className="text-sm text-muted-foreground">
                    No se encontraron enlaces de descarga para esta serie.
                  </p>
                ) : (
                  <ul className="space-y-2 text-sm">
                    {links.map((link, index) => (
                      <li key={link + index}>
                        <a
                          href={link}
                          target="_blank"
                          rel="noopener noreferrer"
                          className="text-primary underline-offset-2 hover:underline"
                        >
                          {link}
                        </a>
                      </li>
                    ))}
                  </ul>
                )}
              </div>
            )}
          </CardContent>
        </Card>
      </div>
    </main>
  )
}

export default App
