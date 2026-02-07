<script>
  const softColors = [
            { bg: '#6366F1' }, // Indigo
            { bg: '#8B5CF6' }, // Violet
            { bg: '#EC4899' }, // Pink
            { bg: '#F59E0B' }, // Amber
            { bg: '#10B981' }, // Emerald
            { bg: '#3B82F6' }, // Blue
            { bg: '#F43F5E' }, // Rose
            { bg: '#06B6D4' }  // Cyan
        ];

        document.querySelectorAll('.random-bg').forEach((card, index) => {
            // Mengambil warna secara berurutan atau acak dari palet terpilih
            const color = softColors[index % softColors.length].bg;
            card.style.backgroundColor = color;
            // Memberikan sedikit glow sesuai warna icon
            card.style.boxShadow = `0 10px 15px -3px ${color}40`;
        });
</script>
