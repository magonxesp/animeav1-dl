/**
 * AnimeAV1-DL - Un programa para extraer enlaces de descarga de animeav1.com
 * Copyright (C) 2025  MagonxESP
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
import { type FormEvent, useEffect, useRef, useState } from "react"

import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { cn } from "@/lib/utils"

type DownloadResponse = {
  links: string[]
}

function App() {
  const [url, setUrl] = useState("")
  const [links, setLinks] = useState<string[] | null>(null)
  const [error, setError] = useState<string | null>(null)
  const [isLoading, setIsLoading] = useState(false)
  const [toast, setToast] = useState<{ message: string; type: "success" | "error" } | null>(null)
  const toastTimeoutRef = useRef<number | null>(null)

  useEffect(() => {
    return () => {
      if (toastTimeoutRef.current) {
        window.clearTimeout(toastTimeoutRef.current)
      }
    }
  }, [])

  const showToast = (message: string, type: "success" | "error") => {
    setToast({ message, type })
    if (toastTimeoutRef.current) {
      window.clearTimeout(toastTimeoutRef.current)
    }

    toastTimeoutRef.current = window.setTimeout(() => {
      setToast(null)
    }, 3000)
  }

  const handleCopyAll = async () => {
    if (!links || links.length === 0) {
      return
    }

    try {
      await navigator.clipboard.writeText(links.join("\n"))
      showToast("Enlaces copiados al portapapeles.", "success")
    } catch {
      showToast("No se pudieron copiar los enlaces.", "error")
    }
  }

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
        throw new Error()
      }

      const data = (await response.json()) as DownloadResponse
      setLinks(data.links ?? [])
    } catch {
      setError("No se pudo obtener la lista de enlaces. Inténtalo de nuevo más tarde.")
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <main className="min-h-screen bg-background">
      <section className="mx-auto flex min-h-screen w-full max-w-3xl flex-col items-center justify-center px-4 py-12">
        <header className="mb-8 text-center">
          <h1 className="text-4xl font-bold tracking-tight text-foreground sm:text-5xl">
            AnimeAV1-DL
          </h1>
          <p className="mt-3 max-w-xl text-balance text-sm text-muted-foreground sm:text-base">
            Introduce la URL de la serie en animeav1.com para obtener los enlaces de descarga desde Mega.
          </p>
        </header>

        <form className="w-full" onSubmit={handleSubmit}>
          <Label htmlFor="url" className="sr-only">
            URL de la serie
          </Label>
          <div className="flex w-full flex-col gap-2 sm:max-w-3xl sm:flex-row sm:items-center sm:gap-3">
            <Input
              id="url"
              type="url"
              placeholder="https://animeav1.com/media/..."
              value={url}
              onChange={(event) => setUrl(event.target.value)}
              required
              className="flex-1"
            />
            <Button type="submit" disabled={isLoading}>
              {isLoading ? "Buscando enlaces..." : "Obtener enlaces"}
            </Button>
          </div>
        </form>

        {error && (
          <p className="mt-4 w-full rounded-xl border border-destructive/40 bg-destructive/10 px-3 py-2 text-sm text-destructive">
            {error}
          </p>
        )}

        {links && (
          <div className="mt-6 w-full space-y-4 rounded-2xl border border-border/70 bg-card/60 p-6 shadow-lg backdrop-blur">
            <div className="flex items-center justify-between gap-4">
              <h2 className="text-lg font-semibold text-foreground">Enlaces de Mega</h2>
              <Button
                type="button"
                variant="outline"
                size="sm"
                onClick={handleCopyAll}
                disabled={!links.length}
              >
                Copiar todo
              </Button>
            </div>

            {links.length === 0 ? (
              <p className="text-sm text-muted-foreground">
                No se encontraron enlaces de descarga para esta serie.
              </p>
            ) : (
              <ul className="space-y-2 text-sm">
                {links.map((link, index) => (
                  <li key={link + index} className="truncate">
                    <a
                      href={link}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="text-primary underline-offset-4 hover:underline"
                    >
                      {link}
                    </a>
                  </li>
                ))}
              </ul>
            )}
          </div>
        )}

        {toast && (
          <div
            className={cn(
              "fixed bottom-6 right-6 z-50 min-w-[220px] rounded-xl border px-4 py-3 text-sm shadow-lg backdrop-blur",
              toast.type === "success"
                ? "border-emerald-300/60 bg-emerald-500/10 text-emerald-900 dark:border-emerald-500/60 dark:bg-emerald-500/15 dark:text-emerald-100"
                : "border-destructive/60 bg-destructive/10 text-destructive"
            )}
          >
            {toast.message}
          </div>
        )}
      </section>
    </main>
  )
}

export default App
