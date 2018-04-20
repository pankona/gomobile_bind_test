
OUTDIR = $(CURDIR)/../out
AAR    = $(OUTDIR)/lib.aar

all:
	mkdir -p $(OUTDIR)
	CGO_ENABLED=1 gomobile bind -v -target=android/arm -o $(AAR)

clean:
	rm -rf $(OUTDIR)/*
