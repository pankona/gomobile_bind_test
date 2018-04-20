
OUTDIR = $(CURDIR)/../out
AAR    = $(OUTDIR)/lib.aar

all:
	mkdir -p $(OUTDIR)
	gomobile bind -v -target=android/arm -o $(AAR)

clean:
	rm -rf $(OUTDIR)/*
