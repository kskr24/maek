import { useMutation, useQueryClient } from '@tanstack/react-query'
import { deleteNote } from '@/queries/services/note-service'
import { useToast } from '@/components/ui/hooks/use-toast'

export const useDeleteNote = ({ onSuccess }: { onSuccess?: () => unknown }) => {
  const qc = useQueryClient()
  const { toast } = useToast()

  return useMutation({
    mutationFn: deleteNote,
    onSuccess: () => {
      if (onSuccess) onSuccess()
    },
    onError: (error) => {
      toast({
        title: 'failed to delete note',
        description: error.toString(),
      })
    },
    onSettled: () => {
      void qc.invalidateQueries({ queryKey: ['notes'] })
    },
  })
}
